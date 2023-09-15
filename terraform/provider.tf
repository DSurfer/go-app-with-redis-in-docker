terraform {
    required_providers {
        vkcs = {
            source = "vk-cs/vkcs"
            version = "~> 0.1.12" 
        }
    }
}

provider "vkcs" {
    # Your user account.
    username = var.user_name

    # The password of the account
    password = var.password

    # The tenant token can be taken from the project Settings tab - > API keys.
    # Project ID will be our token.
    project_id = var.api_key
    
    # Region name
    region = "RegionOne"
    
    auth_url = "https://infra.mail.ru:35357/v3/" 
}