.PHONY: all frontend backend dev

# Start both frontend and backend
dev:
	make backend & make frontend

# Start Vue frontend
frontend:
	cd frontend/vue && npm run dev

# Start Go backend
backend:
	cd cmd && go run main.go

# Install dependencies (first time setup)
install:
	cd frontend/vue && npm install
	go mod tidy