# ------------------ #
# build the frontend #
# ------------------ #

FROM node:18 as frontend_builder
WORKDIR /build
RUN npm i -g pnpm@7.22
COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile
ADD frontend .
RUN pnpm build

# ------------------ #
# build the backend  #
# ------------------ #

FROM golang:1.19-alpine AS backend_builder
RUN apk add --no-cache build-base libwebp libwebp-dev
WORKDIR /tmp/lagersystem_manager
# first only copy files relevant for module downloads (caching)
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY --from=frontend_builder /build/dist /tmp/lagersystem_manager/fake_frontend
# Unit tests
# RUN CGO_ENABLED=0 go test -v
RUN go build -o ./lagersystem_manager .

# ----------- #
# FINAL IMAGE #
# ----------- # 

FROM alpine
RUN apk add ca-certificates libwebp
COPY --from=backend_builder /tmp/lagersystem_manager/lagersystem_manager /app/lagersystem_manager
EXPOSE 3000
CMD ["/app/lagersystem_manager"]
