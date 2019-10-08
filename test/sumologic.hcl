plugin "sumologic" {
    type         = "provider"
    version      = "1.1.0"

    url_template = "https://github.com/SumoLogic/sumologic-terraform-provider/releases/download/v{{ .Version }}/sumologic-terraform-provider_{{ .Version }}_{{ .Os }}_{{ .Arch }}.zip"
    replacements = {
        macOS    = "darwin"
        Linux    = "linux"
        "64-bit" = "amd64"
    }
}
