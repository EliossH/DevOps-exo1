# Get Node image
FROM node:24.9.0-alpine AS run-stage

# Create Working directory
WORKDIR /app

# Copy requirement
COPY package.json package-lock.json ./

# Download requirement
RUN npm install

# Copy application
COPY index.js ./

# Expose port
EXPOSE 8090

# Run the app
CMD ["node", "index.js"]