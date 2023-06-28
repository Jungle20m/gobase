pipeline{
    agent any

    environment {
        DOCKERHUB_CREDENTIALS = credentials('gobase')
    }

    stages {
        stage('Clone') {
            steps {
                git 'https://github.com/Jungle20m/gobase.git'
            }
        }
        stage('Login') {
            steps {
                echo 'Login'
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
}