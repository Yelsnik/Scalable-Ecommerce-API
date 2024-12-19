#BUILDER STAGE
FROM node:23.4-alpine AS builder
WORKDIR /app
COPY  package*.json ./
RUN npm ci
COPY  . .
RUN npm run build

# RUN STAGE
FROM node:23.4-alpine As prod
WORKDIR /app
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package*.json ./
COPY --from=builder /app/node_modules ./node_modules
COPY app.env .

# Set NODE_ENV for production
ENV NODE_ENV=production

# Expose the application's port
EXPOSE 3000
CMD ["node", "dist/src/main.js" ]