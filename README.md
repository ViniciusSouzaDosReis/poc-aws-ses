# Aplicação em Go para envio de e-mail

Este repositório contém uma poc de uma API em Go para disparo de e-mail com suporte para deploy automatizado via GitHub Actions e Terraform.

## Como Rodar a GitHub Action

A GitHub Action configurada no arquivo `deploy.yml` é acionada automaticamente quando há um push para a branch `main` e alterações no diretório `back/`. Para rodar a action manualmente, siga os passos abaixo:

1. **Push para a branch `main`**: A action será acionada automaticamente se houver alterações no diretório `back/`.
2. **Verifique o workflow**: Acesse a aba "Actions" no repositório do GitHub para ver o status do deploy.

### Secrets Necessárias

Para que a GitHub Action funcione corretamente, você precisa configurar as seguintes secrets no repositório:

- `AWS_ACCESS_KEY_ID`: Chave de acesso da AWS.
- `AWS_SECRET_ACCESS_KEY`: Chave secreta da AWS.
- `SENDER_EMAIL`: Email do remetente (usado pelo Terraform).
- `S3_BUCKET_NAME`: Nome do bucket S3 (usado pelo Terraform).
- `DYNAMODB_TABLE_NAME`: Nome da tabela DynamoDB (usado pelo Terraform).

Para adicionar as secrets:

1. Vá para o repositório no GitHub.
2. Clique em "Settings" > "Secrets and variables" > "Actions".
3. Clique em "New repository secret" e adicione cada uma das secrets mencionadas acima.

## Como Fazer o Deploy Manualmente

Se você preferir fazer o deploy manualmente, siga os passos abaixo:

### Pré-requisitos

- Go 1.24.1 instalado.
- Terraform 1.10.5 instalado.
- Credenciais da AWS configuradas no ambiente.

### Passos para o Deploy

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/ViniciusSouzaDosReis/poc-aws-ses
   cd poc-aws-ses
   ```

2. **Navegue até o diretório da aplicação:**
  ```bash
   cd back
   ```

3. **Instale as dependências do Go:**
  ```bash
   go mod download
   ```

4. **Build da aplicação:**
  ```bash
   make build
   ```

5. **Inicialize o Terraform:**
  ```bash
   terraform -chdir=infra init
   ```

6. **Planeje as mudanças do Terraform:**
  ```bash
   terraform -chdir=infra plan -out=tfplan
   ```

7. **Aplique as mudanças do Terraform:**
  ```bash
   terraform -chdir=infra apply -auto-approve
   ```

Durante a aplicação, o Terraform usará as variáveis de ambiente configuradas no arquivo deploy.yml. Certifique-se de que as seguintes variáveis de ambiente estejam configuradas no seu ambiente local:
```bash
export TF_VAR_sender_email=seu-email@example.com
export TF_VAR_s3_bucket_name=meu-bucket-s3
export TF_VAR_dynamodb_table_name=minha-tabela-dynamodb
```

## Rota de Acesso
Após o deploy, a URL base da API será gerada pelo Terraform e estará disponível como um output. O Terraform exibirá a URL no final da execução do comando terraform apply. A URL pode ser acessada da seguinte forma:
```tf
output "base_url" {
  value = aws_api_gateway_stage.example.invoke_url
}
```
Essa URL será exibida no terminal após a execução do terraform apply e será usada para acessar a API.
