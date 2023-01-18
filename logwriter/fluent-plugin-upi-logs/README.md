# fluent-plugin-upi-logs

[Fluentd](https://fluentd.org/) parser plugin to parse [UPI](https://github.com/caraml-dev/universal-prediction-interface) logs proto into JSON format.

## Installation

### RubyGems

```
$ gem install fluent-plugin-upi-logs
```

### Bundler

Add following line to your Gemfile:

```ruby
gem "fluent-plugin-upi-logs"
```

And then execute:

```
$ bundle install
```

## Configuration

```aconf
<parse>
  @type upi_logs
  class_name UPI.Protobuf.Class.Name # eg caraml.upi.v1.PredictionLog
</parse>
```

## Development
`lib/fluent/plugin/parser_upi_logs.rb` is the main parser script that extends fluentd plugin. 

On how to install this custom plugin to be used in fluentd, check the [fluentd installation guide](https://docs.fluentd.org/plugin-development). 
The parser script can be copied over, or simply do a gem/bundler install in the dockerfile

To run test:
```
bundle exec rake test 
```

## LICENSE

[Apache-2.0](LICENSE)
