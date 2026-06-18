import { Link } from "@tanstack/react-router"
import { Card, CardContent } from "@/components/ui/card"
import { Plus } from "lucide-react"

export function AddMonitorCard() {
  return (
    <Link to="/dashboard/add">
      <Card className="bg-[#121212] border-white/10 border-dashed hover:border-emerald-500/50 transition-colors cursor-pointer h-full min-h-[300px] flex flex-col items-center justify-center group">
        <CardContent className="flex flex-col items-center justify-center gap-4 p-8">
          <div className="w-12 h-12 rounded-full bg-white/5 flex items-center justify-center group-hover:bg-emerald-500/10 transition-colors">
            <Plus className="w-6 h-6 text-zinc-400 group-hover:text-emerald-500 transition-colors" />
          </div>
          <p className="text-zinc-400 font-medium">Add New Monitor</p>
        </CardContent>
      </Card>
    </Link>
  )
}
