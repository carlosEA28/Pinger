import { useState } from "react"
import { useNavigate, createFileRoute, Link } from "@tanstack/react-router"
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Button } from "@/components/ui/button"
import { Switch } from "@/components/ui/switch"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { useCreateMonitor } from "@/hooks/use-monitors"
import { Link2 } from "lucide-react"

export const Route = createFileRoute('/dashboard/add')({
  component: AddMonitor,
})

function AddMonitor() {
  const navigate = useNavigate()
  const { mutate: createMonitor, isPending } = useCreateMonitor()

  const [name, setName] = useState("")
  const [url, setUrl] = useState("")
  const [interval, setInterval] = useState("60")
  const [showAdvanced, setShowAdvanced] = useState(false)

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    createMonitor(
      { url, intervalSeconds: parseInt(interval), name },
      {
        onSuccess: () => {
          navigate({ to: "/" })
        },
      }
    )
  }

  return (
    <div className="p-8 flex items-center justify-center">
      <Card className="w-full max-w-md bg-[#121212] border-white/10">
        <CardHeader>
          <div className="w-10 h-10 bg-white/5 rounded-lg flex items-center justify-center mb-4">
            <Link2 className="w-5 h-5 text-emerald-500" />
          </div>
          <CardTitle className="text-2xl">Create New Monitor</CardTitle>
          <CardDescription className="text-zinc-400">
            Configure a new endpoint to track uptime, latency, and SSL certificate validity.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="space-y-2">
              <Label className="text-xs font-bold text-zinc-500 uppercase tracking-wider">Service Name</Label>
              <Input
                placeholder="e.g. Production API"
                className="bg-[#1a1a1a] border-white/10 focus-visible:ring-emerald-500"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            </div>

            <div className="space-y-2">
              <Label className="text-xs font-bold text-zinc-500 uppercase tracking-wider">URL Endpoint</Label>
              <div className="relative">
                <Link2 className="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-zinc-500" />
                <Input
                  placeholder="https://api.service.com/health"
                  className="bg-[#1a1a1a] border-white/10 pl-10 focus-visible:ring-emerald-500"
                  value={url}
                  onChange={(e) => setUrl(e.target.value)}
                  required
                />
              </div>
            </div>

            <div className="space-y-2">
              <Label className="text-xs font-bold text-zinc-500 uppercase tracking-wider">Check Interval</Label>
              <Select value={interval} onValueChange={setInterval}>
                <SelectTrigger className="bg-[#1a1a1a] border-white/10 focus:ring-emerald-500">
                  <SelectValue />
                </SelectTrigger>
                <SelectContent className="bg-[#1a1a1a] border-white/10">
                  <SelectItem value="30">30s</SelectItem>
                  <SelectItem value="60">1m</SelectItem>
                  <SelectItem value="300">5m</SelectItem>
                  <SelectItem value="900">15m</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="space-y-4 pt-4 border-t border-white/10">
              <div className="flex items-center justify-between">
                <div>
                  <h4 className="font-medium">Advanced Settings</h4>
                  <p className="text-xs text-zinc-500">Configure custom headers, timeout, and authentication.</p>
                </div>
                <Switch checked={showAdvanced} onCheckedChange={setShowAdvanced} />
              </div>
            </div>

            <div className="flex gap-4 pt-4">
              <Link to="/" className="flex-1">
                <Button type="button" variant="outline" className="w-full border-white/10 hover:bg-white/5">
                  Cancel
                </Button>
              </Link>
              <Button
                type="submit"
                className="flex-1 bg-emerald-600 hover:bg-emerald-700 text-white"
                disabled={isPending || !url}
              >
                {isPending ? "Creating..." : "Create Monitor"}
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
