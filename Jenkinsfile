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