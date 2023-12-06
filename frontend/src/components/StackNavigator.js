import React from "react";
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { createStackNavigator } from '@react-navigation/stack';
import { NavigationContainer } from '@react-navigation/native';
import { Entypo, AntDesign } from '@expo/vector-icons';

import HomeScreen from '../screens/Common/HomeScreen';
import LoginScreen from '../screens/Common/LoginScreen';

const BottomTabIcons = [
  {
    label: "Home",
    focused: <Entypo name="home" size={24} color="white" />,
    unfocused: <AntDesign name="home" size={24} color="black" />,
    component: HomeScreen,
  },
  {
    label: "Login",
    focused: <Entypo name="login" size={24} color="white" />,
    unfocused: <AntDesign name="login" size={24} color="black" />,
    component: LoginScreen,
  },
];

const StackIcons = [
  {
    name: "Home",
    component: HomeScreen,
  },
  {
    name: "Login",
    component: LoginScreen,
  },
  {
    name: "Main",
    component: BottomTabs,
  },
];

const Tab = createBottomTabNavigator();
const Stack = createStackNavigator();

function BottomTabs() {
  return (
    <Tab.Navigator screenOptions={ScreenOptions}>
      {BottomTabIcons.map((icon, index) => (
        <Tab.Screen
          key={index}
          name={icon.label}
          component={icon.component}
          options={{
            tabBarLabel: icon.label,
            headerShown: false,
            tabBarLabelStyle: { color: "white" },
            tabBarIcon: ({ focused }) =>
              focused ? icon.focused : icon.unfocused,
          }}
        />
      ))}
    </Tab.Navigator>
  )
}

function Navigation() {
  return (
    <NavigationContainer>
      <Stack.Navigator>
        {StackIcons.map((icon, index) => (
          <Stack.Screen
            key={index}
            name={icon.name}
            component={icon.component}
            options={{ headerShown: false }}
          />
        ))}
      </Stack.Navigator>
    </NavigationContainer>
  )
}

const ScreenOptions = {
  tabBarStyle: {
    backgroundColor: "rgba(0, 0, 0, 0.5)",
    position: "absolute",
    bottom: 0,
    right: 0,
    shadowOpacity: 4,
    shadowRadius: 4,
    elevation: 4,
    shadowOffset: {
      width: 0,
      height: -4,
    },
    borderTopWidth: 0,
  },
};

export default Navigation;
