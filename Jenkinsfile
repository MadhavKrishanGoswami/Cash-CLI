
pipeline {
    agent {
        node {
          label 'docker-agent'
        }
      }
    stages {
        stage('Build') {
          steps{
            echo "Build"
            sh '''
              go mod vendor
              go mod tidy
              go build -o Cash-CLI
            '''
          }
        }
        stage('Test'){
          steps{ 
            echo "Testing"
            sh '''
            go test
            '''
          }
        }
        stage('Docker image build'){
          steps{
              echo "Building Docker Image"
              script {
                withDockerRegistry([credentialsId: '4ccd9c79-9f50-4f84-8953-7819845d9b8d', url: 'https://index.docker.io/v1/']) {
                    sh "docker build -t cash-cli ."
                    sh "docker tag cash-cli madhavkrishangoswami/cash-cli:latest"
                }
              }
          }
        }
        stage('Deploying to DockerHub'){
          steps{
            echo "Deploying Docker Image"
            script {
                withDockerRegistry([credentialsId: '4ccd9c79-9f50-4f84-8953-7819845d9b8d', url: 'https://index.docker.io/v1/']) {
                  sh "docker push madhavkrishangoswami/cash-cli:latest"
                }
            }
          }
        }
    }
}

