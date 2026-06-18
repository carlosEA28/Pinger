import { AreaChart, Area, XAxis, YAxis, Tooltip, ResponsiveContainer } from "recharts"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import type { LatencyMetric } from "@/lib/api"

interface LatencyChartProps {
  metrics: LatencyMetric[]
}

export function LatencyChart({ metrics }: LatencyChartProps) {
  const chartData = metrics.map((m) => ({
    time: new Date(m.timestamp).toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" }),
    latency: m.responseTimeMs,
  })).reverse()

  const currentAvg = chartData.length > 0 
    ? (chartData.reduce((acc, curr) => acc + curr.latency, 0) / chartData.length).toFixed(1)
    : "0.0"

  return (
    <Card className="bg-[#121212] border-white/10">
      <CardHeader className="flex flex-row items-start justify-between">
        <div>
          <CardTitle>Latency (ms)</CardTitle>
          <p className="text-xs text-zinc-500 mt-1">Response times across global edge nodes</p>
        </div>
        <div className="text-right">
          <div className="flex items-center gap-2">
            <span className="w-2 h-2 rounded-full bg-emerald-500" />
            <span className="text-2xl font-light">{currentAvg}</span>
            <span className="text-xs text-zinc-500">MS</span>
          </div>
          <p className="text-[10px] text-zinc-500 uppercase font-bold tracking-wider">Current Avg</p>
        </div>
      </CardHeader>
      <CardContent>
        <div className="h-[300px] w-full mt-4">
          <ResponsiveContainer width="100%" height="100%">
            <AreaChart data={chartData} margin={{ top: 10, right: 10, left: 0, bottom: 0 }}>
              <defs>
                <linearGradient id="colorLatency" x1="0" y1="0" x2="0" y2="1">
                  <stop offset="5%" stopColor="#10b981" stopOpacity={0.3} />
                  <stop offset="95%" stopColor="#10b981" stopOpacity={0} />
                </linearGradient>
              </defs>
              <XAxis 
                dataKey="time" 
                axisLine={false} 
                tickLine={false} 
                tick={{ fill: "#71717a", fontSize: 12 }}
                dy={10}
              />
              <YAxis 
                axisLine={false} 
                tickLine={false} 
                tick={{ fill: "#71717a", fontSize: 12 }}
                dx={-10}
              />
              <Tooltip
                contentStyle={{ backgroundColor: "#18181b", border: "1px solid #27272a", borderRadius: "8px" }}
                itemStyle={{ color: "#10b981" }}
              />
              <Area 
                type="monotone" 
                dataKey="latency" 
                stroke="#10b981" 
                strokeWidth={2}
                fillOpacity={1} 
                fill="url(#colorLatency)" 
              />
            </AreaChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  )
}
