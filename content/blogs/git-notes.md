+++
date = "2016-08-11T11:59:31+11:00"
title = "Git Practices"
description="Useful Git commands  & practices for repository management"
+++


## Create a new branch with git 

### Create the branch on your local machine and switch in this branch
        
    $ git checkout -b [name_of_your_new_branch]
		
### Push the branch on git-repository (Github, Bitbucket)

    $ git push origin [name_of_your_new_branch]

When you want to commit something in your branch, be sure to be in your branch.

## List all branches

* You can see all branches created by using :
    
    $ git branch

* Which will show :

    * approval_messages
    master
    master_clean

## Manage branches (Push, Fetch & Merge)

Add a new remote for your branch :

    $ git remote add [name_of_your_remote]

Push changes from your commit into your branch :

    $ git push [name_of_your_new_remote] [name_of_your_branch]

Update your branch when the original branch from official repository has been updated :

    $ git fetch [name_of_your_remote]

Then you need to apply to merge changes, if your branch is derivated from develop you need to do :

    $ git merge [name_of_your_remote]/develop

## Delete branch    

Delete a branch on your local filesystem :

    $ git branch -d [name_of_your_new_branch]

To force the deletion of local branch on your filesystem :

    $ git branch -D [name_of_your_new_branch]

Delete the branch on github :

    $ git push origin :[name_of_your_new_branch]

Compare two branch:

    $ git diff [name_of_branch1]..[name_of_branch2]

---

## Branch merge


### Fast-Forward Merge

Our first example demonstrates a fast-forward merge. The code below creates a new branch, adds two commits to it, then integrates it into the main line with a fast-forward merge.

    ```bash
    ## Start a new feature
    $ git checkout -b new-feature master
    
    ## Edit some files    
    $ git add <file>
    $ git commit -m "Start a feature"
    
    ## Edit some files    
    $ git add <file>
    $ git commit -m "Finish a feature"
    
    ## Merge in the new-feature branch    
    $ git checkout master
    $ git merge new-feature
    $ git branch -d new-feature
    ```

This is a common workflow for short-lived topic branches that are used more as an isolated development than an organizational tool for longer-running features.

Also note that Git should not complain about the git branch -d, since new-feature is now accessible from the master branch.

### 3-Way Merge

The next example is very similar, but requires a 3-way merge because master progresses while the feature is in-progress. This is a common scenario for large features or when several developers are working on a project simultaneously.

    ```bash
    ## Start a new feature
    $ git checkout -b new-feature master

    ## Edit some files
    $ git add <file>
    $ git commit -m "Start a feature"

    ## Edit some files
    $ git add <file>
    $ git commit -m "Finish a feature"

    ## Develop the master branch
    $ git checkout master

    ## Edit some files
    $ git add <file>
    $ git commit -m "Make some super-stable changes to master"

    ## Merge in the new-feature branch
    $ git merge new-feature
    $ git branch -d new-feature
    ```

### Git Rebase Tricks

Problem: You are working with a few experienced devs constantly improving an online shooping site. After you complete the first assignment and ready to commit to master you find someone merged a change that affects or overlaps with the ones you made, and it could lead to bugs in the online-shoppping website. 

Solution: Situations like these are a big example of when you'd want to rebase. Let's say when you created your branch off of the master branch, the master branch was on commit No. 1. Every commit in your branch was put on top of commit #1. When you're ready to merge your branch to master, you find  other developers have some changes and the most recent commit is commit No. 4. **Rebasing is taking all your branch's commits and adding them on top of commit No. 4 instead of commit No. 1.** If you consider commit No. 1 as the "base" of your branch, you're changing that base to the most recent one, commit No. 4. Hence why it's called rebasing!

#### Rebase with conflict 

    ```bash
    git rebase master

    # When there is conflict, the rebase will pause. 
    # You have to manually solve the conflict
    # Add the resolved files to stage and commit it
    git add <Resolved-File>
    git commit 

    # Conttinue the rebase process
    git rebase --continue

    ```

#### Rebase interactively

Rebase to master branch 

    $ git rebase -i master 

Rebase with fixup and autosquash

    
    * fixup (f for short), which acts like “squash”, but discards this commit’s message

    # Commit your changes
    $ git add. 
    $ git commit 


    $ git rebase  -i master
    # 
    $ fixup <COMMIT-ID>




## Git reword 

Reword the last commit message. The command below will open an editor to let you change previous commit message

    $ git commit --amend   

Reword the last commit message and author. 

    $ git commit --amend   --author="Other author <other_author@test.com>"

