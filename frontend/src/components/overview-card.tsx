import { Link } from "@tanstack/react-router"
import { Card, CardContent, CardFooter, CardHeader } from "@/components/ui/card"
import { ArrowRight } from "lucide-react"
import type { Monitor } from "@/lib/api"

interface OverviewCardProps {
  monitor: Monitor
}

function UptimeBar() {
  const blocks = Array.from({ length: 30 }, () => Math.random() > 0.05 ? "up" : "down")

  return (
    <div className="flex gap-[2px] h-4 items-end">
      {blocks.map((status, i) => (
        <div
          key={i}
          className={`flex-1 h-full rounded-[1px] ${
            status === "up" ? "bg-emerald-500" : "bg-red-500"
          }`}
        />
      ))}
    </div>
  )
}

export function OverviewCard({ monitor }: OverviewCardProps) {
  const latency = Math.floor(Math.random() * 100) + 20
  const uptime = (99 + Math.random()).toFixed(2)

  return (
    <Card className="bg-[#121212] border-white/10 hover:border-white/20 transition-colors relative overflow-hidden group">
      <div className="absolute top-4 right-4">
        <div className="w-2 h-2 rounded-full bg-emerald-500" />
      </div>
      <CardHeader className="pb-2">
        <h3 className="font-semibold text-lg">{monitor.url.split("//")[1]?.split("/")[0] || monitor.url}</h3>
        <p className="text-xs text-zinc-500 font-mono truncate max-w-[90%]">{monitor.url}</p>
      </CardHeader>
      <CardContent className="space-y-4">
        <div>
          <p className="text-[10px] font-bold text-zinc-500 uppercase tracking-wider mb-1">Current Latency</p>
          <p className="text-2xl font-light text-emerald-400">{latency} <span className="text-sm text-zinc-500">ms</span></p>
        </div>
        <div>
          <p className="text-[10px] font-bold text-zinc-500 uppercase tracking-wider mb-2">30 Day Uptime</p>
          <div className="flex items-center justify-between mb-1">
            <div className="w-full h-2 bg-zinc-800 rounded-full overflow-hidden mr-4">
              <div className="h-full bg-emerald-500" style={{ width: `${uptime}%` }} />
            </div>
            <span className="text-sm text-emerald-400 font-medium">{uptime}%</span>
          </div>
          <UptimeBar />
        </div>
      </CardContent>
      <CardFooter className="pt-2 border-t border-white/5 flex justify-end">
        <Link
          to="/$monitorId"
          params={{ monitorId: monitor.id }}
          className="text-xs text-zinc-400 hover:text-white flex items-center gap-1 transition-colors"
        >
          View Details <ArrowRight className="w-3 h-3" />
        </Link>
      </CardFooter>
    </Card>
  )
}
