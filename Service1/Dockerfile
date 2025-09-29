# Get Node image
FROM python:3.13.7-alpine3.22 AS run-stage

# Create Working directory
WORKDIR /app

# Copy requirement
COPY requirement.txt ./

# Download requirement
RUN pip install -r requirement.txt

# Copy application
COPY main.py ./

# Expose port
EXPOSE 8070

# Run the app
CMD ["python3", "main.py"]