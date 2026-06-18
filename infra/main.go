package main

import (
	"fmt"
	"os"

	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/netlify/netlify"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "netlify")
		token := os.Getenv("NETLIFY_TOKEN")
		if token == "" {
			return fmt.Errorf("missing NETLIFY_TOKEN environment variable")
		}
		siteName := cfg.Get("siteName")
		if siteName == "" {
			siteName = "pinger-frontend"
		}

		netlifyProvider, err := netlify.NewProvider(ctx, "netlify", &netlify.ProviderArgs{
			Token: pulumi.StringPtr(token),
		})
		if err != nil {
			return err
		}

		deployCommand := `npm ci && NITRO_PRESET=netlify npm run build && npx --yes netlify-cli@latest deploy --prod --no-build --auth "$NETLIFY_AUTH_TOKEN" --site-name "$NETLIFY_SITE_NAME" --dir=dist --functions=.netlify/functions-internal --message "Pulumi deploy"`
		deploy, err := local.NewCommand(ctx, "frontend-deploy", &local.CommandArgs{
			Create: pulumi.String(deployCommand),
			Update: pulumi.String(deployCommand),
			Dir:    pulumi.String("../frontend"),
			Environment: pulumi.StringMap{
				"NETLIFY_AUTH_TOKEN": pulumi.String(token),
				"NETLIFY_SITE_NAME":  pulumi.String(siteName),
			},
			Triggers: pulumi.Array{
				pulumi.NewFileAsset("../frontend/package.json"),
				pulumi.NewFileAsset("../frontend/package-lock.json"),
				pulumi.NewFileAsset("../frontend/vite.config.ts"),
				pulumi.NewFileAsset("../frontend/tsconfig.json"),
				pulumi.NewFileArchive("../frontend/public"),
				pulumi.NewFileArchive("../frontend/src"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{netlifyProvider}))
		if err != nil {
			return err
		}

		ctx.Export("frontendDeploy", deploy.Stdout)
		return nil
	})
}
