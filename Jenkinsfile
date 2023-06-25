pipeline{
    agent any
    stages {
        stage('Clone') {
            steps {
                git 'https://github.com/Jungle20m/gobase.git'
            }
        }
        stage('Test') {
            steps {
                echo 'Test'
            }
        }
        stage('Build') {
            steps {
                withDockerRegistry(credentialsId: 'gobase', url: 'https://index.docker.io/v1/') {
                    sh 'docker build -t vietanhd14cn7/gobase:latest .'
                    sh 'docker push vietanhd14cn7/gobase:latest .'
                }
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploy'
            }
        }
    }
}