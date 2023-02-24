# vite-backend-integration-example

A simple example for Vite backend integration with Go Echo.

## dev

`pnpm run dev` then `cd src-backend && go run . live` , try to edit `src/App.tsx` and check hot reload on `127.0.0.1:1323`

## prod

`pnpm run build && cd src-backend && go build .` , then execute `backend.exe` , static files has been embedded in exe.
