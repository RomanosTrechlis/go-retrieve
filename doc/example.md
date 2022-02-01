In this example, we use **retemp** to retrieve typescript code templates for a brand-new React application.

This is not a follow along tutorial, but an actual use case from our daily work life.

### Initializing the project

We execute the following command to create an empty React typescript project.

```bash
yarn create react-app retemp-example --template typescript
```

After the execution of the command we have the following project structure.

![](images/init_project_structure.png)

### Retrieving template

We already have initialized **retemp** and created the necessary registry files.

This is the template list configured with **retemp**.

```bash
retemp template list
```

![](images/template_list.png)

We execute the following command to retrieve the *breadcrumb* template component.

```bash
retemp template breadcrumb -d src/components
```

The flag *-d* puts the retrieved component to the specific directory.

As we see from the template list, the *breadcrumb* component has an
after retrieval command to execute. The following screenshot displays 
the output of the command.

![](images/react-router-dom.png)

The new project structure is as follows:

![](images/end_project_structure.png)

We can, clearly, see the *breadcrumb* component in the specified directory
containing the corresponding typescript files.

This component is now ready to use in the React application.