import { createFileRoute, useNavigate, Link } from '@tanstack/react-router'
import { Button } from "@/components/ui/button"
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog"
import { ArrowLeft, Trash2 } from "lucide-react"
import { useMonitorMetrics, useMonitors, useDeleteMonitor } from "@/hooks/use-monitors"
import { LatencyChart } from "@/components/latency-chart"
import { MetricCard } from "@/components/metric-card"
import { SystemLogs } from "@/components/system-logs"
import { UptimeBadge } from "@/components/uptime-badge"
import { Globe, Wifi, Clock } from "lucide-react"

export const Route = createFileRoute('/$monitorId')({
  component: MonitorDetail,
})

function MonitorDetail() {
  const navigate = useNavigate()
  const { monitorId } = Route.useParams()

  const { data: monitors } = useMonitors()
  const { data: metrics, isLoading } = useMonitorMetrics(monitorId)
  const { mutate: deleteMonitor, isPending } = useDeleteMonitor()

  const monitor = monitors?.find((m) => m.id === monitorId)

  const handleDelete = () => {
    deleteMonitor(monitorId, {
      onSuccess: () => {
          navigate({ to: "/" })
      },
    })
  }

  if (!monitor) {
    return (
      <div className="p-8 flex items-center justify-center">
        <div className="text-center">
          <h2 className="text-2xl font-bold mb-4">Monitor not found</h2>
            <Link to="/">
            <Button variant="outline" className="border-white/10">
              <ArrowLeft className="w-4 h-4 mr-2" /> Back to Dashboard
            </Button>
          </Link>
        </div>
      </div>
    )
  }

  const latestMetric = metrics?.[0]
  const uptime = metrics?.length
    ? ((metrics.filter(m => m.isUp).length / metrics.length) * 100).toFixed(2)
    : "100.00"

  return (
    <div className="p-8">
      <div className="max-w-6xl mx-auto space-y-8">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-6">
          <Link to="/">
              <Button variant="outline" size="icon" className="border-white/10 hover:bg-white/5">
                <ArrowLeft className="w-4 h-4" />
              </Button>
            </Link>
            <div>
              <h1 className="text-2xl font-bold flex items-center gap-3">
                {monitor.url.split("//")[1]?.split("/")[0] || monitor.url}
                <span className="text-sm text-zinc-500 font-mono font-normal">{monitor.url}</span>
              </h1>
            </div>
          </div>
          <div className="flex items-center gap-4">
            <Tabs defaultValue="7d">
              <TabsList className="bg-[#121212] border border-white/10">
                <TabsTrigger value="24h" className="data-[state=active]:bg-white/10">24h</TabsTrigger>
                <TabsTrigger value="7d" className="data-[state=active]:bg-white/10">7d</TabsTrigger>
                <TabsTrigger value="30d" className="data-[state=active]:bg-white/10">30d</TabsTrigger>
              </TabsList>
            </Tabs>
            <Dialog>
              <DialogTrigger asChild>
                <Button variant="outline" size="icon" className="border-red-500/30 text-red-500 hover:bg-red-500/10 hover:text-red-400">
                  <Trash2 className="w-4 h-4" />
                </Button>
              </DialogTrigger>
              <DialogContent className="bg-[#1a1a1a] border-white/10">
                <DialogHeader>
                  <DialogTitle>Delete Monitor</DialogTitle>
                  <DialogDescription className="text-zinc-400">
                    Are you sure you want to delete this monitor? This action cannot be undone.
                  </DialogDescription>
                </DialogHeader>
                <DialogFooter className="gap-2">
                  <DialogTrigger asChild>
                    <Button variant="outline" className="border-white/10 hover:bg-white/5">
                      Cancel
                    </Button>
                  </DialogTrigger>
                  <Button
                    variant="destructive"
                    onClick={handleDelete}
                    disabled={isPending}
                  >
                    {isPending ? "Deleting..." : "Delete"}
                  </Button>
                </DialogFooter>
              </DialogContent>
            </Dialog>
          </div>
        </div>

        {isLoading ? (
          <div className="space-y-6">
            <div className="h-[400px] bg-[#121212] rounded-xl animate-pulse border border-white/5" />
            <div className="grid grid-cols-3 gap-6">
              {[1, 2, 3].map(i => (
                <div key={i} className="h-[140px] bg-[#121212] rounded-xl animate-pulse border border-white/5" />
              ))}
            </div>
          </div>
        ) : (
          <>
            <LatencyChart metrics={metrics || []} />

            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
              <MetricCard
                icon={<Globe className="w-4 h-4" />}
                label="DNS Lookup"
                value={latestMetric?.dnsLookupMs || 0}
                unit="ms"
                description="Domain name to IP resolution time."
              />
              <MetricCard
                icon={<Wifi className="w-4 h-4" />}
                label="TCP Connect"
                value={latestMetric?.tcpConnectMs || 0}
                unit="ms"
                description="Three-way handshake duration."
              />
              <MetricCard
                icon={<Clock className="w-4 h-4" />}
                label="TTFB"
                value={latestMetric?.ttfbMs || 0}
                unit="ms"
                description="Time to First Byte received."
              />
            </div>

            <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
              <div className="lg:col-span-2">
                <SystemLogs metrics={metrics || []} />
              </div>
              <div>
                <UptimeBadge uptime={parseFloat(uptime)} isOperational={monitor.isActive} />
              </div>
            </div>
          </>
        )}
      </div>
    </div>
  )
}
