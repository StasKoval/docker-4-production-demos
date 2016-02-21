require 'rack'
require 'rack/server'
require 'socket'

class HelloWorldApp
  def self.private_ips
    Socket.ip_address_list.select { |intf| intf.ipv4_private? }.map(&:ip_address)
  end

  def self.call(env)
    [200, {}, [
      <<-EOF
<!DOCTYPE html>
<html>
<head>
<title>Hello World!</title>
<style>body { width: 35em; margin: 0 auto; font-family: Tahoma, Verdana, Arial, sans-serif; }</style>
</head>
<body>
<h1>Hello world!</h1>
<p>This is an example of a very very simple app.</p>
<p>Server's (private) IP: #{private_ips.last}</p>
</body>
</html>
EOF
    ]]
  end
end

run HelloWorldApp
