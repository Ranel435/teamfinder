# Build stage
FROM node:20-alpine as builder

WORKDIR /app

COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Production stage
FROM node:20-alpine

WORKDIR /app

# Copy built assets from builder stage
COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json .
COPY --from=builder /app/node_modules ./node_modules

EXPOSE 3000

CMD ["node", "build"]
