import React from 'react'

interface DashboardCardProps {
  title: string
  value?: number
}

const DashboardCard = ({ title, value = 0 }: DashboardCardProps) => {
  return (
    <div className="bg-white shadow-md p-6 rounded-md">
      <h2 className="text-xl font-bold">{title}</h2>
      <p className="text-3xl">{value}</p>
    </div>
  )
}

export default DashboardCard
