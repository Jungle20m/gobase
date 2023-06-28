pipeline{
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = credentials('github')
    }
    stages {
        stage('Clone') {
            steps {
                git 'https://github.com/Jungle20m/gobase.git'
            }
        }
        stage('Login') {
            steps {
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
            }
        }
        stage('Build') {
            steps {
                echo 'Build'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploy'
            }
        }
    }
    post {
        always {
            sh 'docker logout'
        }
    }
}