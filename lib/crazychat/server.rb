require 'crazychat/error'
require 'sinatra/base'
require 'sinatra/reloader'
require 'json'

module CrazyChat
  class Server < Sinatra::Base
    register Sinatra::Reloader

    post '/' do
      begin
        body = request.body.read
        json = JSON.parse(body, symbolize_names: true)
        username, address, message = json.values_at(:username, :returnAddress, :message)

        settings.chat_client.handle_message(username, address, message)
        204 # OK
      rescue CrazyChatError => e
        [400, e.message]
      end
    end
  end
end
