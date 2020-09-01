from django.urls import path, include
from . import views
from django.contrib.auth.views import LoginView

urlpatterns = [
    path('', views.index, name='index'),
    path('index/', views.index, name='index'),
    
    path('add/', views.addTodo, name='add'),
    path('complete/<todo_id>/', views.completeTodo, name='complete'),
    path('update/<todo_id>/', views.updateTodo, name='update'),
    path('deletecomplete/', views.deletecomplete, name='deletecomplete'),
    path('deleteall/', views.deleteall, name='deleteall'),

    path('login/', views.loginpage, name='login'),
    path('register/', views.registerpage, name='register'),
    path('logout/', views.logoutpage, name='logout'),
]