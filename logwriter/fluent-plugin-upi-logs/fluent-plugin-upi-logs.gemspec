# frozen_string_literal: true

Gem::Specification.new do |spec|
  spec.name    = 'fluent-plugin-upi-logs'
  spec.version = '0.0.0.dev.1'
  spec.authors = ['Caraml Dev']
  spec.email   = ['caraml-dev@gojek.com']
  spec.license = 'Apache-2.0'

  spec.summary       = 'Fluent parser plugin for UPI logs into JSON'
  spec.description   = 'Fluentd parser custom plugin that can parse UPI logs (PredictionLog and RouterLog
   - https://github.com/caraml-dev/universal-prediction-interface) into json'
  spec.homepage      = 'https://github.com/caraml-dev/timber'

  _test_files = Dir['test/**/*.rb']
  spec.files = Dir['lib/**/*.rb']
  spec.require_paths = ['lib']
  spec.required_ruby_version = '~> 3.1'

  spec.add_development_dependency 'bundler', '~> 2.3.26'
  spec.add_development_dependency 'rake', '~> 13.0.6'
  spec.add_development_dependency 'rubocop', '~> 1.43.0'
  spec.add_development_dependency 'test-unit', '~> 3.5.7'
  spec.add_runtime_dependency 'caraml-upi-protos', ['~> 0.0.0']
  spec.add_runtime_dependency 'fluentd', ['>= 0.14.10', '< 2']
  spec.add_runtime_dependency 'google-protobuf', ['~> 3.12']
  spec.add_runtime_dependency 'ruby-protocol-buffers', ['~> 1.5.0']
end
