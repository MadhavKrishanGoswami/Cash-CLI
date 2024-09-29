pipeline {
    agent any
    stages {
        stage('Build') {
          echo "Build"
          sh '''
          go build -o Cash-CLI
          '''
        }
        stage('Test'){
          echo "Testing"
          sh '''
          go test
          '''
        }
        stage('Dokcer image build'){
            echo "Building Docker Image"
            scripts{
              withDockerRegistry(credentialsId: '4ccd9c79-9f50-4f84-8953-7819845d9b8d') {
                  sh "docker build -t cash-cli ."
                  sh "docker tag cash-cli madhavkrishangoswami/cash-cli:latest"
                  sh "docker push madhavkrishangoswami/cash-cli:latest"
              }
            }
        }
        stage('Deploying to DockerHub'){
            echo "Deploying Docker Image"
            scripts{
              withDockerRegistry(credentialsId: '4ccd9c79-9f50-4f84-8953-7819845d9b8d') {
                  sh "docker push madhavkrishangoswami/cash-cli:latest"
              }
            }
        }
    }
}
