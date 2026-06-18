import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { cn } from "@/lib/utils"

interface UptimeBadgeProps {
  uptime: number
  isOperational: boolean
}

export function UptimeBadge({ uptime, isOperational }: UptimeBadgeProps) {
  const blocks = Array.from({ length: 7 }, () => Math.random() > 0.05 ? "up" : "down")

  return (
    <Card className="bg-[#121212] border-white/10">
      <CardHeader>
        <CardTitle className="text-sm font-medium">Uptime</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="flex gap-1 h-4">
          {blocks.map((status, i) => (
            <div
              key={i}
              className={cn(
                "flex-1 h-full rounded-[2px]",
                status === "up" ? "bg-emerald-500" : "bg-zinc-800"
              )}
            />
          ))}
        </div>
        <div className="flex items-end justify-between">
          <span className="text-xs text-zinc-500 uppercase tracking-widest font-bold">30 Days</span>
          <span className="text-xs text-zinc-500 uppercase tracking-widest font-bold">Today</span>
        </div>
        <div className="pt-4 border-t border-white/5">
          <div className="flex items-baseline gap-1">
            <span className="text-4xl font-light">{uptime}</span>
            <span className="text-xl text-zinc-500">%</span>
          </div>
          <div className="flex items-center gap-2 mt-2">
            <span className={cn(
              "w-2 h-2 rounded-full",
              isOperational ? "bg-emerald-500" : "bg-red-500"
            )} />
            <span className={cn(
              "text-xs font-bold uppercase tracking-wider",
              isOperational ? "text-emerald-500" : "text-red-500"
            )}>
              {isOperational ? "Operational" : "Degraded"}
            </span>
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
