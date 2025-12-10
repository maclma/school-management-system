# Remove tracked local artifacts (school.db, logs/, .env) and commit .gitignore
Set-Location -Path "$PSScriptRoot\.."

if (git ls-files | Select-String -Pattern '^school.db$') {
    Write-Host "Untracking school.db"
    git rm --cached school.db
} else {
    Write-Host "school.db not tracked"
}

if (git ls-files | Select-String -Pattern '^\.env$') {
    Write-Host "Untracking .env"
    git rm --cached .env
} else {
    Write-Host ".env not tracked"
}

if (git ls-files | Select-String -Pattern '^logs/') {
    Write-Host "Untracking logs/"
    git rm -r --cached logs
} else {
    Write-Host "logs not tracked"
}

# Stage .gitignore and commit if changes
git add .gitignore
if ((git status --porcelain) -ne '') {
    git commit -m "Remove local DB, logs and .env from repo; add to .gitignore"
    Write-Host "Committed changes."
} else {
    Write-Host "Nothing to commit"
}
