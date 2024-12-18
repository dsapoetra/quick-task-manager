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
    <div className="min-h-screen bg-gray-50">
        <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        {/* User Profile Card */}
        <div className="mb-6 bg-white rounded-lg shadow">
          <div className="px-4 py-5 sm:px-6 border-b border-gray-200">
            <h3 className="text-lg font-medium text-gray-900">User Profile</h3>
          </div>
          <div className="px-4 py-5 sm:p-6">
            <div className="flex items-center space-x-6">
              <div className="flex-shrink-0">
                <div className="h-24 w-24 rounded-full bg-indigo-600 flex items-center justify-center">
                  <span className="text-3xl font-bold text-white">
                    {user?.username?.charAt(0).toUpperCase()}
                  </span>
                </div>
              </div>
              <div>
                <h4 className="text-xl font-medium text-gray-900">{user?.username}</h4>
                <p className="text-sm text-gray-500">{user?.email}</p>
                <div className="mt-2 flex space-x-4">
                  <button className="text-sm text-indigo-600 hover:text-indigo-900">
                    Edit Profile
                  </button>
                  <button className="text-sm text-indigo-600 hover:text-indigo-900">
                    Change Password
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      {/* Navbar */}
      <nav className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16">
            <div className="flex items-center">
              <h1 className="text-xl font-bold text-gray-900">TaskManager</h1>
            </div>
            <div className="flex items-center space-x-4">
              <span className="text-sm text-gray-500">Welcome, {user?.username}</span>
              <button 
                onClick={() => { /* Add logout logic */ }}
                className="text-sm text-red-600 hover:text-red-800"
              >
                Logout
              </button>
            </div>
          </div>
        </div>
      </nav>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        {/* Stats Overview */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-sm font-medium text-gray-500">Total Tasks</h3>
            <p className="text-3xl font-bold text-gray-900">{tasks.length}</p>
          </div>
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-sm font-medium text-gray-500">In Progress</h3>
            <p className="text-3xl font-bold text-indigo-600">
              {tasks.filter(task => task.status === 'IN_PROGRESS').length}
            </p>
          </div>
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-sm font-medium text-gray-500">DONE</h3>
            <p className="text-3xl font-bold text-green-600">
              {tasks.filter(task => task.status === 'DONE').length}
            </p>
          </div>
        </div>

        {/* Tasks List */}
        <div className="bg-white shadow rounded-lg">
          <div className="px-4 py-5 sm:px-6 flex justify-between items-center">
            <h2 className="text-lg font-medium text-gray-900">Your Tasks</h2>
            <button className="bg-indigo-600 text-white px-4 py-2 rounded-md text-sm hover:bg-indigo-700">
              New Task
            </button>
          </div>
          <ul className="divide-y divide-gray-200">
            {tasks.map((task) => (
              <li key={task.id} className="px-4 py-4 sm:px-6 hover:bg-gray-50">
                <div className="flex items-center justify-between">
                  <div className="flex-1 min-w-0">
                    <h3 className="text-sm font-medium text-gray-900 truncate">
                      {task.title}
                    </h3>
                    <p className="mt-1 text-sm text-gray-500">
                      {task.description}
                    </p>
                  </div>
                  <div className="flex items-center space-x-4">
                    <span className={`px-2 py-1 text-xs rounded-full ${
                      task.priority === 'HIGH' 
                        ? 'bg-red-100 text-red-800'
                        : task.priority === 'MEDIUM'
                        ? 'bg-yellow-100 text-yellow-800'
                        : 'bg-green-100 text-green-800'
                    }`}>
                      {task.priority}
                    </span>
                    <span className={`px-2 py-1 text-xs rounded-full ${
                      task.status === 'DONE'
                        ? 'bg-green-100 text-green-800'
                        : task.status === 'IN_PROGRESS'
                        ? 'bg-blue-100 text-blue-800'
                        : 'bg-gray-100 text-gray-800'
                    }`}>
                      {task.status}
                    </span>
                  </div>
                </div>
              </li>
            ))}
          </ul>
        </div>
      </main>
    </div>
    </div>
  );
};

export default Dashboard;