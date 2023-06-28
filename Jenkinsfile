pipeline{
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub')
    }
    stages {
        stage('Clone') {
            steps {
                git 'https://github.com/Jungle20m/gobase.git'
            }
        }
        stage('Build') {
            steps {
                echo 'docker build -t vietanhd14cn7/gobase:latest'
            }
        }
        stage('Login') {
            steps {
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
            }
        }

        stage('Push') {
            steps {
                echo 'docker push vietanhd14cn7/gobase:latest'
            }
        }
    }
    post {
        always {
            sh 'docker logout'
        }
    }
}