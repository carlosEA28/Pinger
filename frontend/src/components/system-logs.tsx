import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import type { LatencyMetric } from "@/lib/api"
import { cn } from "@/lib/utils"

interface SystemLogsProps {
  metrics: LatencyMetric[]
}

export function SystemLogs({ metrics }: SystemLogsProps) {
  const recentMetrics = metrics.slice(0, 5)

  return (
    <Card className="bg-[#121212] border-white/10">
      <CardHeader className="flex flex-row items-center justify-between">
        <CardTitle>System Logs</CardTitle>
        <Badge variant="outline" className="border-emerald-500/50 text-emerald-500 bg-emerald-500/10 text-[10px] px-2 py-0.5">
          Live
        </Badge>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          {recentMetrics.length > 0 ? recentMetrics.map((metric) => (
            <div key={metric.id} className="flex items-center justify-between text-sm font-mono">
              <div className="flex items-center gap-4">
                <span className={cn(
                  "font-bold",
                  metric.statusCode >= 200 && metric.statusCode < 300 ? "text-emerald-500" : "text-red-500"
                )}>
                  {metric.statusCode}
                </span>
                <span className={cn(
                  "text-xs px-1.5 py-0.5 rounded",
                  metric.statusCode >= 200 && metric.statusCode < 300 
                    ? "bg-emerald-500/10 text-emerald-500" 
                    : "bg-red-500/10 text-red-500"
                )}>
                  {metric.statusCode >= 200 && metric.statusCode < 300 ? "OK" : "ERR"}
                </span>
                <span className="text-zinc-400">GET /v1/health</span>
              </div>
              <span className="text-zinc-500">
                {new Date(metric.timestamp).toLocaleTimeString([], { hour: "2-digit", minute: "2-digit", second: "2-digit" })}
              </span>
            </div>
          )) : (
            <div className="text-center text-zinc-500 py-8">No logs available</div>
          )}
        </div>
      </CardContent>
    </Card>
  )
}
