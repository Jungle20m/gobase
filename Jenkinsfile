pipeline{
    agent any
    stages {
        stage('Clone') {
            steps {
                git 'https://github.com/Jungle20m/gobase.git'
                checkout scmGit(branches: [[name: '*/12-feature-integrate-cicd']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/Jungle20m/gobase.git']])
            }
        }
    }
}