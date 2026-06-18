import { Card, CardContent, CardHeader } from "@/components/ui/card"

interface MetricCardProps {
  label: string
  value: string | number
  unit: string
  description: string
  icon: React.ReactNode
}

export function MetricCard({ label, value, unit, description, icon }: MetricCardProps) {
  return (
    <Card className="bg-[#121212] border-white/10">
      <CardHeader className="flex flex-row items-center gap-2 pb-2">
        <div className="text-emerald-500">{icon}</div>
        <h3 className="text-xs font-bold text-zinc-500 uppercase tracking-wider">{label}</h3>
      </CardHeader>
      <CardContent>
        <div className="flex items-baseline gap-1 mb-2">
          <span className="text-3xl font-light">{value}</span>
          <span className="text-sm text-zinc-500">{unit}</span>
        </div>
        <p className="text-xs text-zinc-500">{description}</p>
      </CardContent>
    </Card>
  )
}
