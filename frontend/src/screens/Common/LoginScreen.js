import React from "react";
import { StyleSheet, ScrollView, Text } from "react-native";
import { LinearGradient } from "expo-linear-gradient";

export default function LoginScreen() {
  return (
    <LinearGradient 
      colors={["#4c669f", "#3b5998", "#192f6a"]}
    >
      <ScrollView style={styles.container}>
        <Text>Login Screen</Text>
      </ScrollView>
    </LinearGradient>
  )
}

const styles = StyleSheet.create({
  container: {
    marginTop: 50,
  }
});