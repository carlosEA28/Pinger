import { createFileRoute } from '@tanstack/react-router'
import { OverviewCard } from "@/components/overview-card"
import { AddMonitorCard } from "@/components/add-monitor-card"
import { useMonitors } from "@/hooks/use-monitors"
import { Button } from "@/components/ui/button"
import { RefreshCcw } from "lucide-react"

export const Route = createFileRoute('/')({
  component: DashboardIndex,
})

function DashboardIndex() {
  const { data: monitors, isLoading, refetch } = useMonitors()

  return (
    <div className="p-8">
      <div className="max-w-7xl mx-auto space-y-8">
        <div className="flex items-end justify-between">
          <div>
            <h1 className="text-3xl font-bold mb-2">System Overview</h1>
            <p className="text-zinc-500">Real-time monitoring of global services infrastructure.</p>
          </div>
          <div className="flex gap-4">
            <Button
              variant="outline"
              className="border-white/10 hover:bg-white/5"
              onClick={() => refetch()}
            >
              <RefreshCcw className="w-4 h-4 mr-2" /> Refresh
            </Button>
          </div>
        </div>

        {isLoading ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {[1, 2, 3].map((i) => (
              <div key={i} className="h-[400px] bg-[#121212] rounded-xl animate-pulse border border-white/5" />
            ))}
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {monitors?.map((monitor) => (
              <OverviewCard key={monitor.id} monitor={monitor} />
            ))}
            <AddMonitorCard />
          </div>
        )}
      </div>
    </div>
  )
}
