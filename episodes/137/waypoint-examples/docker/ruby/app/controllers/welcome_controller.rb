class WelcomeController < ApplicationController

  # GET /welcome
  def index
    @host = request.host_with_port
  end

end
