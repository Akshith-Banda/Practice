from django import forms
from .models import TodoModel
from django.contrib.auth.forms import UserCreationForm
from django.contrib.auth.models import User

class TodoForm(forms.ModelForm):
    class Meta:
        model = TodoModel
        fields = ['text']
        widgets = {
            'text' : forms.TextInput(attrs={'placeholder' : 'enter things to do...', 'class' : 'form-control'})
        }

class CreateUserForm(UserCreationForm):
    class Meta:
        model = User
        fields = ['username', 'email']
        widgets = {
            'username' : forms.TextInput(attrs={'placeholder' : 'username...', 'class' : 'form-control'}),
            'email' : forms.EmailInput(attrs={'placeholder' : 'email...', 'class' : 'form-control'}),
            #'password1' : forms.PasswordInput(attrs={'placeholder' : 'password...', 'class' : 'form-control'}),
            #'password2' : forms.PasswordInput(attrs={'placeholder' : 'password confirm...', 'class' : 'form-control'}),
        }
    def __init__(self, *args, **kwargs):
        super(CreateUserForm, self).__init__(*args, **kwargs)
        self.fields['password1'].widget = forms.PasswordInput(attrs={'class': 'form-control', 'placeholder': 'Password..'})
        self.fields['password2'].widget = forms.PasswordInput(attrs={'class': 'form-control', 'placeholder': 'Password confirmation..'})