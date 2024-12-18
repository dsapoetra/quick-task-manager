export interface UserData {
  id: string;
  username: string;
  email: string;
}

export interface DecodedToken {
  user: UserData;
  exp: number;
  iat: number;
}



export interface Task {
  id: number;
  title: string;
  description: string;
  status: string;
  assignee_id: number;
  assigner_id: number;
  priority: number;
  created_at: string;
  updated_at: string;
}