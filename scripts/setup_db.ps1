param(
    [string]$DatabaseUrl = $env:DATABASE_URL,
    [string]$DbName = $env:DB_NAME,
    [string]$DbUser = $env:DB_USER,
    [string]$DbPassword = $env:DB_PASSWORD,
    [string]$DbHost = $env:DB_HOST,
    [string]$DbPort = $env:DB_PORT
)

if ([string]::IsNullOrWhiteSpace($DbName))
{ $DbName = "ayma" 
}
if ([string]::IsNullOrWhiteSpace($DbUser))
{ $DbUser = "postgres" 
}
if ([string]::IsNullOrWhiteSpace($DbPassword))
{ $DbPassword = "postgres" 
}
if ([string]::IsNullOrWhiteSpace($DbHost))
{ $DbHost = "localhost" 
}
if ([string]::IsNullOrWhiteSpace($DbPort))
{ $DbPort = "5432" 
}

if ([string]::IsNullOrWhiteSpace($DatabaseUrl))
{
    $DatabaseUrl = "postgres://$DbUser`:$DbPassword@$DbHost`:$DbPort/$DbName?sslmode=disable"
}

Write-Host "Using DATABASE_URL=$DatabaseUrl"

createdb $DbName
psql $DatabaseUrl -f migrations/001_init.sql
psql $DatabaseUrl -f migrations/002_seed.sql
