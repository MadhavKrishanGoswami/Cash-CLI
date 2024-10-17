<p align="center">
  <img src="https://i.imgur.com/ddhPSQ4.png" height="64">
  <h3 align="center">Cash-CLI</h3>
  <p align="center">Convert Currency Rates directly from your Terminal</p>

![demo](https://github.com/user-attachments/assets/74b564bb-5e8d-4887-94bf-28c212a00407)

This is a command-line interface (CLI) currency converter built with Go. The project was created as a learning exercise to apply DevOps knowledge and best practices.

## Features

- Convert between different currencies using up-to-date exchange rates
- Simple and intuitive command-line interface (TUI)

## DevOps Practices Applied

- [x] Version Control (Git)
- [x] Containerization (Docker)
- [x] CI/CD Pipeline (GitHub Actions)
- [ ] Infrastructure as Code (Terraform)
- [ ] Automated Testing (Go testing package)
- [ ] Monitoring and Logging (Prometheus & Grafana)
- [ ] Secrets Management (HashiCorp Vault)
- [ ] Continuous Deployment (ArgoCD)
- [ ] Load Balancing
- [ ] High Availability Setup

## Installation

### Prerequisites

- Go 1.16 or higher
- Docker (optional, for running containerized version)

### From Source

1. Clone the repository:
   ```
   git clone https://github.com/MadhavKrishanGoswami/Cash-CLI.git
   ```

2. Navigate to the project directory:
   ```
   cd Cash-CLI
   ```

3. Build the application:
   ```
   go build -o Cash-CLI
   ```

### Using Docker

1. Pull the Docker image:
   ```
   docker pull madhavkrishangoswami/cash-cli
   ```
2. Run the Docker image:
   ```
   docker run --rm -it --name cash-cli madhavkrishangoswami/cash-cli 

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
