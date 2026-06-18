import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query"
import { api, type Monitor, type LatencyMetric } from "@/lib/api"

export function useMonitors() {
  return useQuery<Monitor[]>({
    queryKey: ["monitors"],
    queryFn: api.getMonitors,
  })
}

export function useMonitorMetrics(id: string) {
  return useQuery<LatencyMetric[]>({
    queryKey: ["monitors", id, "metrics"],
    queryFn: () => api.getMonitorMetrics(id),
    enabled: !!id,
  })
}

export function useCreateMonitor() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: api.createMonitor,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["monitors"] })
    },
  })
}

export function usePingMonitor() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: api.pingMonitor,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["monitors"] })
    },
  })
}

export function useDeleteMonitor() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: api.deleteMonitor,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["monitors"] })
    },
  })
}
