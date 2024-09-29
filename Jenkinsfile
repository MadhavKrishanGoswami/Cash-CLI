pipeline {
    agent any
    stages {
        stage('ls') {
            steps {
                sh '''
                ls 
                '''
              }
          }
        stage('Build') {
            steps {
                echo "Building.."
                sh '''
                docker build -t cash-cli .
                '''
            }
        }
        stage('Test') {
            steps {
                echo "Testing.."
                sh '''
                echo "Test's here"
                '''
            }
        }
        stage('Log in to DockerHub') {
            steps {
                echo 'Logging In....'
                sh '''
                docker login -u $DOCKERHUB_USERNAME -P $DOCKERHUB_PASS
                '''
            }
        }
        stage('Deliver') {
            steps {
                echo 'Pushing To DockerHub....'
                sh '''
                docker push madhavkrishangoswami/cash-cli:latest
                '''
            }
        }

    }
}
