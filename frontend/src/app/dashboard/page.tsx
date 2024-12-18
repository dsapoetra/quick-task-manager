'use client'

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';  // Change this line
import { isTokenValid, getUserFromToken } from '../utils/auth.utils';
import { UserData, Task } from '../types/auth.types';

const API_URL = process.env.NEXT_PUBLIC_API_URL;

const Dashboard: React.FC = () => {
    const router = useRouter();  // Change this line
    const [user, setUser] = useState<UserData | null>(null);
    const [tasks, setTasks] = useState<Task[]>([]);

  useEffect(() => {
    if (!isTokenValid()) {
      router.push('/api/auth/login');
      return;
    }

    const userData = getUserFromToken();
    setUser(userData);
  }, [router]);

  // Optional: Refresh user data from API
  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await fetch(`${API_URL}/api/auth/profile`, {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('jwt_token')}`
          }
        });
        if (response.ok) {
          const userData = await response.json();
          setUser(userData);
        }
      } catch (error) {
        console.error('Failed to fetch user data:', error);
      }
    };

    const fetchTasks = async () => {
      const response = await fetch(`${API_URL}/api/task/assigner`, {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('jwt_token')}`
        }
      });
      const tasks = await response.json();
      setTasks(tasks);
    };  

    fetchUserData();
    fetchTasks();
  }, []);


  if (!user) {
    return <div>Loading...</div>;
  }

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Welcome to Dashboard</h1>
      <div className="bg-white p-4 rounded-lg shadow">
        <p className="text-lg">Welcome back, {user.username}!</p>
        <p className="text-gray-600">Email: {user.email}</p>
      </div>
    <div>
      <h2>Your Tasks</h2>
      {/* Add a list of tasks here */}
      {tasks.map((task) => (
        <div key={task.id}>
          <h3>{task.title}</h3>
          <p>{task.description}</p>
          <p>{task.status}</p>
          <p>{task.priority}</p>
          <p>{task.created_at}</p>
          <p>{task.updated_at}</p>
          <p>{task.assignee_id}</p>
          <p>{task.assigner_id}</p>
        </div>
      ))}
    </div>
    </div>
  );
};

export default Dashboard;