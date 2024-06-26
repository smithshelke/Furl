<div align="center">

```
    ______          _ 
   |  ____|        | |
   | |__ _   _ _ __| |
   |  __| | | | '__| |
   | |  | |_| | |  | |
   |_|   \__,_|_|  |_|
 
```
</div>
                    
<div align="center">
Furl - A generic automation framework
</div>


Goal
--------------------------
All companies in the world can expose their client facing business workflows using a dot furl file. All users should be able to make any workflow automations using the publicly available workflows.

Requirements
--------------------------

*  create blocks
*  store configurations of blocks
*  block can be of many types such as a network block, wait block, cron block
*  make workflows from blocks
*  each workflow should be executable
*  each block can output the json
*  each block can take an input
*  should store users
*  user authentication
*  user has workflow/s
*  workflow can be made of multiple workflows
*  worksflows are owned by a user
*  workflows and blocks have description
*  workflows can be shared to multiple users, (workflows are duplicated)
*  should list most popular/used workflows 
*  integrate an ai layer to automatically create workflows, using description and metainfo of workflows and blocks
*  import and export dot furl files
*  provide a few built in most commonly used workflows such as authentication etc 


Todo
--------------------------
* create block executor
* create apis to edit block
* create apis to edit workflow
* create apis to edit user
* create workflow runs
* connect mysql for logs
* connect redis
* create WorflowRunEngine
* create graph traversers
* import workflows
* share workflows


In progress
--------------------------
* model and design network block executor


Completed
--------------------------
* create apis to create user
* create apis to list user
* create apis to create block
* create apis to create workflow
* create apis to assign blocks/workflows to user
* create apis to connect blocks/workflows

