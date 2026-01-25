group "default" {
  targets = [ "client", "server" ] 
}

target "client" {
  context = "./client" 
  dockerfile = "Dockerfile"
  tags = [ "ghcr.io/yassentials/game-tic-tac-toe/client:latest" ]
}

target "server" {
  context = "./server"
  dockerfile = "Dockerfile"
  tags = [ "ghcr.io/yassentials/game-tic-tac-toe/server:latest" ]
}