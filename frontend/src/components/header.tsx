import { Link } from "@tanstack/react-router"
import { Button } from "@/components/ui/button"
import { useMonitors } from "@/hooks/use-monitors"

export function HeaderComponent() {
  const { data: monitors } = useMonitors()
  const onlineCount = monitors?.filter((m) => m.isActive).length ?? 0
  const totalCount = monitors?.length ?? 0

  return (
    <header className="flex items-center justify-between px-6 py-4 border-b border-white/10 bg-[#0a0a0a]">
      <div className="flex items-center gap-8">
        <Link to="/" className="text-xl font-bold tracking-tight">
          Uptime<span className="text-emerald-500">Pinger</span>
        </Link>
        <nav className="flex items-center gap-6 text-sm text-zinc-400">
          <Link
            to="/"
            className="hover:text-white transition-colors"
            activeProps={{ className: "text-white font-medium" }}
          >
            Dashboard
          </Link>
          <Link
            to="/"
            className="hover:text-white transition-colors"
          >
            Incidents
          </Link>

        </nav>
      </div>

      <div className="flex items-center gap-4">
        <div className="text-sm text-emerald-400 font-medium">
          <span className="text-white">{onlineCount}/{totalCount}</span> Systems Online
        </div>
        <Button 
          variant="outline" 
          className="border-white/20 hover:bg-white/10 text-white"
          onClick={() => window.location.href = '/dashboard/add'}
        >
          Add Monitor
        </Button>
      </div>
    </header>
  )
}
