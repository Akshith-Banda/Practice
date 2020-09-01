from django.shortcuts import render, redirect
from django.views.decorators.http import require_POST
from .models import TodoModel
from .forms import TodoForm, CreateUserForm
import datetime
from django.utils import timezone
from django.contrib.auth import authenticate, login, logout
from django.contrib.auth.decorators import login_required

def index(request):
    date_today = timezone.now()
    if request.user.is_authenticated:
        todo_list = TodoModel.objects.filter(user=request.user).all()
    else:
        todo_list = ''
    todo_form = TodoForm()
    context= {'todo_list' : todo_list, 'todo_form' : todo_form, 'date_today' : date_today}
    return render(request, 'todo/index.html', context)

@login_required(login_url='index')
@require_POST
def addTodo(request):
     if request.method == 'POST':
         todo_form = TodoForm(request.POST)
         if todo_form.is_valid():
            instance = todo_form.save(commit=False)
            instance.user = request.user
            instance.save()
         return redirect('index')

@login_required(login_url='index')
def completeTodo(request, todo_id):
    todo_item = TodoModel.objects.get(pk=todo_id)
    todo_item.completed = True
    todo_item.save()
    return redirect('index')

@login_required(login_url='index')
def updateTodo(request, todo_id):
    date_today = timezone.now()
    todo_item = TodoModel.objects.get(pk=todo_id)
    todo_form = TodoForm(instance=todo_item)
    todo_list = TodoModel.objects.order_by('id')
    if request.method == 'POST':
        todo_form = TodoForm(request.POST, instance=todo_item)
        if todo_form.is_valid():
            todo_form.save()
            return redirect('index')
    context= {'todo_list' : todo_list, 'todo_form' : todo_form, 'date_today' : date_today}
    return render(request, 'todo/update.html', context)

@login_required(login_url='index')
def deletecomplete(request):
    TodoModel.objects.filter(completed__exact=True).delete()
    return redirect('index')

@login_required(login_url='index')
def deleteall(request):
    TodoModel.objects.all().delete()
    return redirect('index')

def loginpage(request):
    if request.user.is_authenticated:
        return redirect('index')
    else:
        if request.method == 'POST':
            username = request.POST.get('username')
            password = request.POST.get('password')
            user = authenticate(username=username, password=password)
            if user is not None:
                login(request, user)
                return redirect('index')
        context = {}
        return render(request, 'registration/login.html', context)

def registerpage(request):
    if request.user.is_authenticated:
        return redirect('index')
    else:
        form = CreateUserForm()
        if request.method == 'POST':
            form = CreateUserForm(request.POST)
            if form.is_valid():
                form.save()
                return redirect('login')
        context = {'form' : form}
        return render(request, 'registration/register.html', context)

def logoutpage(request):
    logout(request)
    return redirect('index')









