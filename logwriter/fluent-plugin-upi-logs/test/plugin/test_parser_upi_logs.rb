# frozen_string_literal: true

require 'helper'
require 'fluent/plugin/parser_upi_logs'
require 'google/protobuf/well_known_types'

require 'caraml/upi/v1/prediction_log_pb'

class UpiParserTest < Test::Unit::TestCase
  setup do
    Fluent::Test.setup
  end

  def create_driver(conf)
    Fluent::Test::Driver::Parser.new(Fluent::Plugin::UpiParser).configure(conf)
  end

  sub_test_case 'configure' do
    test 'empty conf' do
      err = assert_raise(Fluent::ConfigError) { create_driver('') }
      assert_equal err.message, "'class_name' parameter is required"
    end

    test 'invalid conf' do
      err = assert_raise(Fluent::ConfigError) { create_driver(%(class_name "non existence proto")) }
      assert_equal err.message, "unrecognised class_name 'non existence proto'"
    end

    test 'valid conf' do
      create_driver(%(class_name "caraml.upi.v1.PredictionLog"))
      create_driver(%(class_name "caraml.upi.v1.RouterLog"))
    end
  end

  sub_test_case 'plugin will parse text' do
    test 'parse ok' do
      d = create_driver(%(class_name "caraml.upi.v1.PredictionLog"))
      msg = ::Caraml::Upi::V1::PredictionLog.new(
        prediction_id: '1',
        model_name: '123',
        input: ::Caraml::Upi::V1::ModelInput.new(
          'features_table' => Google::Protobuf::Struct.from_hash({
                                                                   'subkey' => 123.456789,
                                                                   'subkey2' => false,
                                                                   'subkey3' => 'hello'
                                                                 })
        )
      )
      binary = ::Caraml::Upi::V1::PredictionLog.encode(msg)
      expected = { 'prediction_id' => '1',
                   'model_name' => '123',
                   'input' =>
                      { 'features_table' =>
                         { 'subkey3' => 'hello', 'subkey2' => false, 'subkey' => 123.456789 } } }
      d.instance.parse(binary) do |time, record|
        assert_equal(expected, record)
        assert_not_nil(time)
      end
    end
  end
end
