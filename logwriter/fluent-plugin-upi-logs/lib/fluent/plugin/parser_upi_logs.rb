# frozen_string_literal: true

require 'json'

require 'fluent/plugin/parser'

require 'caraml/upi/v1/prediction_log_pb'
require 'caraml/upi/v1/router_log_pb'

module Fluent
  module Plugin
    # Parser for UPI which supports prediction and router logs
    class UpiParser < Fluent::Plugin::Parser
      # Parser name should be identical to "parser_#{name}" to follow fluentd plugin convention
      Fluent::Plugin.register_parser('upi_logs', self)

      config_param :class_name, :string

      def configure(conf)
        super
        if Google::Protobuf::DescriptorPool.generated_pool.lookup(@class_name).nil?
          raise Fluent::ConfigError, "unrecognised class_name '#{class_name}'"
        end

        # Lookup will only work for proto that are declared by the require statements
        @protobuf_descriptor = Google::Protobuf::DescriptorPool.generated_pool.lookup(@class_name).msgclass
      end

      def parse(binary)
        record = @protobuf_descriptor.decode(binary)
        time = Fluent::EventTime.now
        # Record are returned in json format. 'to_h' is not used directly as it will retain all fields with zero value
        # '.to_json' will omit empty fields from the proto and then parsed with JSON lib
        puts time
        yield time, JSON.parse(record.to_json({ preserve_proto_fieldnames: true }))
      end
    end
  end
end
