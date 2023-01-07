# my-tech-blog

# URL
https://mytechblog.link/


# 使用技術
- go 1.19.1
  - echo v4.9.1
  - sqlx v1.3.5
- MySQL 8.0.28
- AWS
  - VPC
  - EC2
  - RDS
  - Route53
  - ELB
  - ECS
  - ECR
- Docker
- Github Actions CI/CD

# AWS構成図
![AWS](https://user-images.githubusercontent.com/37237733/211139670-6e7df8f2-71fd-4cfc-a66f-d7fe678513a8.png)

# Github Actions CI/CD
- Githubへのpush時に、buildが自動で実行されます。
- mainブランチへのpushでは、buildが成功した場合、ECSへの自動デプロイが実行されます

# 機能一覧
- 新規記事作成
- 記事編集
- 記事閲覧
- 記事削除
