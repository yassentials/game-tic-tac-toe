group "default" {
  targets = [ "client", "server" ] 
}

target "client" {
  context = "./client" 
  dockerfile = "Dockerfile"
  tags = [ "ghcr.io/up9t/game-tic-tac-toe/frontend:latest" ]
}

target "server" {
  context = "./server"
  dockerfile = "Dockerfile"
  tags = [ "ghcr.io/up9t/game-tic-tac-toe/backend:latest" ]
}
