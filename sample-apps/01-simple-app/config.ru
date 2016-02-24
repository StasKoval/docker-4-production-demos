require 'rack'
require 'rack/server'
require 'socket'

class HelloWorldApp
  def self.call(env)
    [200, {}, [
      <<-EOF
<!DOCTYPE html>
<html>
<head>
<title>Hello World!</title>
<style>
body { width: 35em; margin: 0 auto; font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Hello world! Version: 1</h1>
<p><strong>Hostname:</strong> #{Socket.gethostname}</p>
</body>
</html>
EOF
    ]]
  end
end

run HelloWorldApp
