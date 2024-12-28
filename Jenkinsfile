pipeline {
    agent any

    environment {
        DOCKER_IMAGE_SERVER = 'thahphuog/devopsasm'
        DOCKER_IMAGE_CLIENT = 'thahphuog/client'
        DOCKER_TAG = '1.0.0'
        TELEGRAM_BOT_TOKEN = '7908085505:AAEy0dz1yrVesOaFmZ1s5qWlvslKWekBi_k'
        TELEGRAM_CHAT_ID = '-1002403309943'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'master', url: 'https://github.com/Thanh-Phuog/asm_devops.git'
            }
        }


           stage('Build Client Docker Image') {
            steps {
                script {
                    echo 'Building Client Docker image'
                    docker.build("${DOCKER_IMAGE_CLIENT}:${DOCKER_TAG}", './book_mana_client')
                }
            }
        }

     
        stage('Build Server Docker Image') {
            steps {
                script {
                    echo 'Building Server Docker image'
                    docker.build("${DOCKER_IMAGE_SERVER}:${DOCKER_TAG}", './book_mana_server')
                }
            }
        }


        stage('Push to Docker Hub') {
            steps {
                script {
                    // Push Server image
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE_SERVER}:${DOCKER_TAG}").push()
                    }

                    // Push Client image
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE_CLIENT}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy with Docker Compose') {
            steps {
                script {
                    echo 'Deploying using Docker Compose...'
                    sh 'docker-compose -f docker-compose.yml up -d'
                }
            }
        }
    }
    post {
        always {
            cleanWs()
        }

        success {
            sendTelegramMessage("✅ Build #${BUILD_NUMBER} was successful! ✅")
        }

        failure {
            sendTelegramMessage("❌ Build #${BUILD_NUMBER} failed. ❌")
        }
    }
}

def sendTelegramMessage(String message) {
    sh """
    curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage \
    -d chat_id=${TELEGRAM_CHAT_ID} \
    -d text="${message}"
    """
}
