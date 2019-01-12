+++
tags =  ["itext"]
date = "2016-08-11T11:59:31+11:00"
title = "Git Notes"
description="Common Git Commands & Tricks"
+++


Create a new branch with git and manage branches

* Create the branch on your local machine and switch in this branch :
        
        $ git checkout -b [name_of_your_new_branch]
		
* Push the branch on github :

        $ git push origin [name_of_your_new_branch]

When you want to commit something in your branch, be sure to be in your branch.


You can see all branches created by using :
    
    $ git branch

Which will show :

    * approval_messages
      master
      master_clean

Add a new remote for your branch :

    $ git remote add [name_of_your_remote]

Push changes from your commit into your branch :

    $ git push [name_of_your_new_remote] [name_of_your_branch]

Update your branch when the original branch from official repository has been updated :

    $ git fetch [name_of_your_remote]

Then you need to apply to merge changes, if your branch is derivated from develop you need to do :

    $ git merge [name_of_your_remote]/develop

Delete a branch on your local filesystem :

    $ git branch -d [name_of_your_new_branch]

To force the deletion of local branch on your filesystem :

    $ git branch -D [name_of_your_new_branch]

Delete the branch on github :

    $ git push origin :[name_of_your_new_branch]

Compare two branch:

    $ git diff [name_of_branch1]..[name_of_branch2]

---

Branch merge

Example

Fast-Forward Merge

Our first example demonstrates a fast-forward merge. The code below creates a new branch, adds two commits to it, then integrates it into the main line with a fast-forward merge.

    ```bash
    #### Start a new feature
    $ git checkout -b new-feature master
    
    #### Edit some files    
    $ git add <file>
    $ git commit -m "Start a feature"
    
    #### Edit some files    
    $ git add <file>
    $ git commit -m "Finish a feature"
    
    #### Merge in the new-feature branch    
    $ git checkout master
    $ git merge new-feature
    $ git branch -d new-feature
    ```

This is a common workflow for short-lived topic branches that are used more as an isolated development than an organizational tool for longer-running features.

Also note that Git should not complain about the git branch -d, since new-feature is now accessible from the master branch.

3-Way Merge

The next example is very similar, but requires a 3-way merge because master progresses while the feature is in-progress. This is a common scenario for large features or when several developers are working on a project simultaneously.

    ```bash
    #### Start a new feature
    $ git checkout -b new-feature master

    #### Edit some files
    $ git add <file>
    $ git commit -m "Start a feature"

    #### Edit some files
    $ git add <file>
    $ git commit -m "Finish a feature"

    #### Develop the master branch
    $ git checkout master

    #### Edit some files
    $ git add <file>
    $ git commit -m "Make some super-stable changes to master"

    #### Merge in the new-feature branch
    $ git merge new-feature
    $ git branch -d new-feature
    ```