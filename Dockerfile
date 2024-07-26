FROM swaggerapi/swagger-codegen-cli-v3

ENTRYPOINT ["java", "-jar", "/opt/swagger-codegen-cli/swagger-codegen-cli.jar", "generate", "-i", "/swagger/api/swagger.yaml", "-l", "go", "-o", "/swagger"]