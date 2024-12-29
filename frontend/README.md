This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

# Quick Task Manager Frontend

Welcome to the **Quick Task Manager Frontend** repository! This project is part of the Quick Task Manager application and serves as the user interface layer, enabling users to interact with the task management system efficiently.

## Table of Contents

- [About](#about)
- [Features](#features)
- [Technologies](#technologies)
- [Setup](#setup)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## About

The Quick Task Manager frontend is designed to provide a seamless and intuitive experience for managing tasks. This project communicates with the backend to fetch and display tasks, as well as to create, update, and delete tasks.

## Features

- View a list of tasks
- Add new tasks
- Edit existing tasks
- Delete tasks
- Drag and drop tasks for easy reordering
- Responsive design for desktop and mobile

## Technologies

- **Framework:** React.js
- **State Management:** React Context API or Redux (based on implementation)
- **Styling:** CSS or a library such as TailwindCSS (if applicable)
- **API Communication:** Axios or Fetch API
- **Drag and Drop:** @hello-pangea/dnd

## Setup

To set up and run the project locally, follow these steps:

### Prerequisites

Ensure you have the following installed on your machine:

- Node.js (v16 or later)
- npm or yarn

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/dsapoetra/quick-task-manager.git
   cd quick-task-manager/frontend
   ```

2. Install dependencies:

   ```bash
   npm install
   # or
   yarn install
   ```

### Environment Variables

Create a `.env` file in the project root and add the following variables:

```
NEXT_PUBLIC_API_URL=http://localhost:8080
```

Adjust the `NEXT_PUBLIC_API_URL` to match your backend API URL.

## Usage

### Development Server

To start the development server:

```bash
npm start
# or
yarn start
```

The application will run at [http://localhost:3000](http://localhost:3000).

### Build for Production

To build the application for production:

```bash
npm run build
# or
yarn build
```

The build files will be generated in the `build/` directory.

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature-name`)
3. Commit your changes (`git commit -m 'Add feature'`)
4. Push to the branch (`git push origin feature-name`)
5. Open a pull request

## License

This project is licensed under the MIT License. See the [LICENSE](../LICENSE) file for details.

---

Feel free to reach out for any questions or suggestions. Happy coding!

